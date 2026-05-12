#!/usr/bin/env python3
"""Sync LeetCode solutions metadata from Go source files to Google Sheets.

Appends NEW problems and updates Solve Date for existing problems.
Never modifies or deletes other existing data. Matches the existing sheet format:
  Problem Number | Difficulty | Is Competent | Solve Date | Next Solve Date | Topics | Topics 2
"""

import glob
import json
import os
import re
import subprocess
import sys

import gspread
from google.oauth2.service_account import Credentials

HEADER_PATTERN = re.compile(
    r"//\s*LeetCode\s+#(\d+)\s+\((Easy|Medium|Hard)\):\s+(https://leetcode\.com/problems/([^/\s]+))"
)

SCOPES = [
    "https://www.googleapis.com/auth/spreadsheets",
    "https://www.googleapis.com/auth/drive",
]


def get_credentials():
    """Load Google service account credentials from env var (JSON string or file path)."""
    creds_json = os.environ.get("GOOGLE_SHEETS_CREDENTIALS_JSON")
    creds_file = os.environ.get("GOOGLE_SHEETS_CREDENTIALS_FILE")

    if creds_json:
        info = json.loads(creds_json)
        return Credentials.from_service_account_info(info, scopes=SCOPES)
    elif creds_file:
        return Credentials.from_service_account_file(creds_file, scopes=SCOPES)
    else:
        print("Error: Set GOOGLE_SHEETS_CREDENTIALS_JSON or GOOGLE_SHEETS_CREDENTIALS_FILE")
        sys.exit(1)


def get_sheet_id():
    """Get the Google Sheet ID from env var."""
    sheet_id = os.environ.get("GOOGLE_SHEET_ID")
    if not sheet_id:
        print("Error: Set GOOGLE_SHEET_ID environment variable")
        sys.exit(1)
    return sheet_id


def find_solution_files(repo_root):
    """Find all .go solution files (excluding test files)."""
    pattern = os.path.join(repo_root, "**", "*.go")
    files = glob.glob(pattern, recursive=True)
    return [f for f in files if not f.endswith("_test.go")]


def parse_header(filepath):
    """Extract problem metadata from the standardized comment header."""
    with open(filepath, "r") as f:
        content = f.read()

    match = HEADER_PATTERN.search(content)
    if not match:
        return None

    number = int(match.group(1))
    difficulty = match.group(2)

    return {
        "number": number,
        "difficulty": difficulty,
    }


def get_topic(filepath, repo_root):
    """Derive topic from the parent directory name."""
    rel = os.path.relpath(filepath, repo_root)
    parts = rel.split(os.sep)
    if len(parts) >= 2:
        return parts[0].replace("_", " ").title()
    return "Unknown"


def get_date_solved(filepath, repo_root):
    """Get the date the file was first added to git, formatted as M/D/YYYY."""
    try:
        result = subprocess.run(
            ["git", "log", "--diff-filter=A", "--follow", "--format=%aI", "--", filepath],
            cwd=repo_root,
            capture_output=True,
            text=True,
            timeout=30,
        )
        if result.returncode == 0 and result.stdout.strip():
            lines = result.stdout.strip().splitlines()
            iso_date = lines[-1].strip()[:10]
            # Convert YYYY-MM-DD to M/D/YYYY to match existing sheet format
            parts = iso_date.split("-")
            if len(parts) == 3:
                return f"{int(parts[1])}/{int(parts[2])}/{parts[0]}"
    except (subprocess.TimeoutExpired, FileNotFoundError):
        pass
    return ""


def collect_problems(repo_root):
    """Collect metadata for all solution files."""
    files = find_solution_files(repo_root)
    problems = []

    for filepath in files:
        meta = parse_header(filepath)
        if meta is None:
            print(f"  Skipping (no valid header): {filepath}")
            continue

        topic = get_topic(filepath, repo_root)
        date_solved = get_date_solved(filepath, repo_root)

        problems.append({
            "number": meta["number"],
            "difficulty": meta["difficulty"],
            "topic": topic,
            "date_solved": date_solved,
        })

    problems.sort(key=lambda p: p["number"])
    return problems


def get_existing_problem_rows(worksheet):
    """Read column A to get all problem numbers and their row indices."""
    col_a = worksheet.col_values(1)
    existing = {}
    for i, val in enumerate(col_a[1:], start=2):  # Skip header, rows are 1-indexed
        val = val.strip()
        if val.isdigit():
            existing[int(val)] = i
    return existing


def sync_to_sheet(problems, creds, sheet_id):
    """Append new problems and update Solve Date for existing ones."""
    client = gspread.authorize(creds)
    spreadsheet = client.open_by_key(sheet_id)

    # Use the first available worksheet (handles renamed/reordered sheets)
    worksheet = spreadsheet.worksheets()[0]

    # Get problem numbers already in the sheet (mapped to row index)
    existing_rows = get_existing_problem_rows(worksheet)
    print(f"  {len(existing_rows)} problems already in sheet")

    # Update Solve Date (column D) for existing problems
    updates = []
    for p in problems:
        if p["number"] in existing_rows and p["date_solved"]:
            row = existing_rows[p["number"]]
            updates.append({"range": f"D{row}", "values": [[p["date_solved"]]]})

    if updates:
        worksheet.batch_update(updates, value_input_option="USER_ENTERED")
        print(f"  Updated Solve Date for {len(updates)} existing problems")

    # Filter to only new problems
    new_problems = [p for p in problems if p["number"] not in existing_rows]

    if not new_problems:
        print("  No new problems to add")
        return

    # Build rows matching existing format:
    # Problem Number | Difficulty | Is Competent | Solve Date | Next Solve Date | Topics | Topics 2
    rows = []
    for p in new_problems:
        rows.append([
            str(p["number"]),   # Problem Number
            p["difficulty"],   # Difficulty
            "",                # Is Competent (Brandon fills in)
            p["date_solved"],  # Solve Date
            "",                # Next Solve Date (Brandon fills in)
            p["topic"],        # Topics
            "",                # Topics 2 (Brandon fills in)
        ])

    # Find the actual last row with data and write directly below it
    all_values = worksheet.get_all_values()
    next_row = len(all_values) + 1

    # Ensure the sheet has enough rows
    if worksheet.row_count < next_row + len(rows) - 1:
        worksheet.add_rows(next_row + len(rows) - 1 - worksheet.row_count)

    cell_range = f"A{next_row}:G{next_row + len(rows) - 1}"
    worksheet.update(values=rows, range_name=cell_range, value_input_option="USER_ENTERED")

    print(f"  Appended {len(new_problems)} new problems:")
    for p in new_problems:
        print(f"    #{p['number']} ({p['difficulty']}) - {p['topic']}")


def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    repo_root = os.path.dirname(script_dir)

    print("Collecting problem metadata from repo...")
    problems = collect_problems(repo_root)
    print(f"  Found {len(problems)} problems in repo")

    if not problems:
        print("  No problems found. Check that solution files have standardized headers.")
        sys.exit(1)

    for p in problems:
        print(f"  #{p['number']} ({p['difficulty']}) - {p['topic']} - {p['date_solved']}")

    print("\nSyncing to Google Sheets...")
    creds = get_credentials()
    sheet_id = get_sheet_id()
    sync_to_sheet(problems, creds, sheet_id)

    print("Done!")


if __name__ == "__main__":
    main()

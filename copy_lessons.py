#!/usr/bin/env python3
"""
Copy Lessons Script

This script copies directories from the 'source/lessons' directory to the 'work/lessons' directory.
If a directory already exists in the work directory, it will be skipped to preserve existing work.

Usage:
    python copy_lessons.py [--preview]

Options:
    --preview    Show what directories would be copied without actually copying them
"""

import os
import sys
import argparse
import shutil
from pathlib import Path
from typing import List, Tuple


def get_directories_to_copy(source_dir: Path, work_dir: Path) -> Tuple[List[str], List[str]]:
    """
    Analyze source and work directories to determine what needs to be copied.
    
    Args:
        source_dir: Path to source/lessons directory
        work_dir: Path to work/lessons directory
    
    Returns:
        Tuple of (directories_to_copy, directories_to_skip)
    """
    if not source_dir.exists():
        print(f"Error: Source directory '{source_dir}' does not exist.")
        return [], []
    
    # Ensure work directory exists
    work_dir.mkdir(parents=True, exist_ok=True)
    
    # Get all directories in source
    source_dirs = [d.name for d in source_dir.iterdir() if d.is_dir()]
    
    # Get all directories in work
    work_dirs = [d.name for d in work_dir.iterdir() if d.is_dir()]
    
    # Determine what to copy and what to skip
    to_copy = []
    to_skip = []
    
    for dir_name in source_dirs:
        if dir_name in work_dirs:
            to_skip.append(dir_name)
        else:
            to_copy.append(dir_name)
    
    return sorted(to_copy), sorted(to_skip)


def copy_directory(source_path: Path, dest_path: Path) -> bool:
    """
    Copy a directory from source to destination.
    
    Args:
        source_path: Source directory path
        dest_path: Destination directory path
    
    Returns:
        True if successful, False otherwise
    """
    try:
        shutil.copytree(source_path, dest_path)
        return True
    except Exception as e:
        print(f"Error copying {source_path} to {dest_path}: {e}")
        return False


def main():
    parser = argparse.ArgumentParser(
        description="Copy lesson directories from source to work directory",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
    python copy_lessons.py --preview    # Show what would be copied
    python copy_lessons.py              # Actually copy the directories
        """
    )
    
    parser.add_argument(
        "--preview",
        action="store_true",
        help="Show what directories would be copied without actually copying them"
    )
    
    args = parser.parse_args()
    
    # Define paths
    script_dir = Path(__file__).parent
    source_lessons_dir = script_dir / "source" / "lessons"
    work_lessons_dir = script_dir / "work" / "lessons"
    
    print(f"Source directory: {source_lessons_dir}")
    print(f"Work directory: {work_lessons_dir}")
    print()
    
    # Analyze directories
    to_copy, to_skip = get_directories_to_copy(source_lessons_dir, work_lessons_dir)
    
    if args.preview:
        print("=== PREVIEW MODE ===")
        print("This is what would happen if you run the script without --preview:")
        print()
    
    # Report what will be copied
    if to_copy:
        if args.preview:
            print(f"📁 Directories that WOULD BE COPIED ({len(to_copy)}):")
        else:
            print(f"📁 Copying directories ({len(to_copy)}):")
        
        for dir_name in to_copy:
            source_path = source_lessons_dir / dir_name
            dest_path = work_lessons_dir / dir_name
            
            if args.preview:
                print(f"   ✅ {dir_name}")
                print(f"      From: {source_path}")
                print(f"      To:   {dest_path}")
            else:
                print(f"   📋 Copying {dir_name}...", end=" ")
                if copy_directory(source_path, dest_path):
                    print("✅ Success")
                else:
                    print("❌ Failed")
        print()
    else:
        print("📁 No new directories to copy.")
        print()
    
    # Report what will be skipped
    if to_skip:
        if args.preview:
            print(f"⏭️  Directories that WOULD BE SKIPPED ({len(to_skip)}):")
        else:
            print(f"⏭️  Skipping existing directories ({len(to_skip)}):")
        
        for dir_name in to_skip:
            print(f"   ⚠️  {dir_name} (already exists in work directory)")
        print()
    
    # Summary
    if args.preview:
        print("=== SUMMARY ===")
        print(f"Total directories in source: {len(to_copy) + len(to_skip)}")
        print(f"Would copy: {len(to_copy)}")
        print(f"Would skip: {len(to_skip)}")
        print()
        print("To actually perform the copy operation, run:")
        print("    python copy_lessons.py")
    else:
        if to_copy:
            print("=== COPY OPERATION COMPLETE ===")
            print(f"Successfully processed {len(to_copy)} directories.")
            if to_skip:
                print(f"Skipped {len(to_skip)} existing directories.")
        else:
            print("=== NO ACTION NEEDED ===")
            print("All source directories already exist in the work directory.")


if __name__ == "__main__":
    main()


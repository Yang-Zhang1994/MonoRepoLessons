# CS 5500 - Foundations of Software Engineering
## Course Materials Repository

Welcome to the CS 5500 Foundations of Software Engineering course materials repository! This repository contains all the exercises, examples, and resources you'll need throughout the semester.

## 🍴 Getting Started: Fork This Repository

**Each student should FORK this repository to their own GitHub account.** This allows you to:
- Collaborate with teammates through pair programming and group learning
- Push your work to your own repository for progress tracking
- Pull updates from the main course repository
- Keep your work safe and version-controlled
- Share and compare solutions with teammates

**Note**: While each student maintains their own fork, these lessons are collaborative learning opportunities. Teams are encouraged to pair program, work together, and learn from each other!

### Initial Setup:
1. Fork this repository on GitHub
2. Clone YOUR fork to your local machine:
   ```bash
   git clone https://github.com/YOUR_USERNAME/MonoRepoLessons.git
   ```
3. Add the original repository as an upstream remote:
   ```bash
   git remote add upstream https://github.com/5500-2025-foundations-yvr/MonoRepoLessons.git
   ```

## ⚠️ Important Notice

**This repository will be updated regularly throughout the term.** New lessons, exercises, and materials will be added as we progress through the course. You'll need to pull these updates from the upstream repository.

## 📁 Repository Structure

```
├── source/
│   └── lessons/          # Course materials and exercises (READ-ONLY)
│       ├── lesson-01/
│       ├── lesson-02/
│       └── ...
├── work/
│   └── lessons/          # Your working directory for exercises
│       ├── lesson-01/
│       ├── lesson-02/
│       └── ...
└── README.md
```

### Directory Explanation

- **`source/lessons/`**: Contains all official course materials, exercises, starter code, and examples. **DO NOT MODIFY FILES IN THIS DIRECTORY.**
- **`work/lessons/`**: Your personal workspace where you should copy and work on exercises.

## 🚨 Critical Workflow Instructions

### Before Starting Any Exercise:

1. **Always copy files from `source/lessons/` to `work/lessons/`** before making any modifications
2. **Never work directly in the `source/lessons/` directory**
3. **Your work in `work/lessons/` will be preserved during repository updates**

### Example Workflow:

```bash
# When starting lesson-01
cp -r source/lessons/lesson-01/ work/lessons/lesson-01/
cd work/lessons/lesson-01/
# Now you can safely modify files and complete the exercise
```

## 🔄 Keeping Your Repository Updated

To get the latest course materials from the main repository:

```bash
# Fetch updates from the main course repository
git fetch upstream

# Merge updates into your main branch
git checkout main
git merge upstream/main

# Push the updates to your fork
git push origin main
```

**Workflow Summary:**
1. **Work in the `work/` directory** for all your exercises
2. **Commit and push your work** to your fork regularly
3. **Pull updates** from upstream when new materials are released
4. **Resolve any merge conflicts** if they occur (rare, since you work in `work/` and updates come to `source/`)

**Important**: Repository updates may overwrite or remove files in the `source/` directory. This is why you must always work in the `work/` directory.

## 📋 Each Exercise is Self-Contained

Each lesson in this repository is designed to be self-contained with:
- Clear instructions and objectives
- All necessary starter files
- Dependencies and setup instructions
- Expected deliverables
- Comprehensive guidance for LLM collaboration

**Note**: Some lessons focus on advanced terminal UI development using libraries like **blessed** for creating rich, interactive command-line applications. These lessons emphasize proper software architecture patterns like MVC (Model-View-Controller) and collaborative development with AI tools.

You don't need to complete exercises in any particular order unless explicitly stated in the lesson instructions.

## 🛡️ Protecting Your Work

With the fork workflow, your work is protected in two ways:
1. **Your fork**: All your commits are safely stored in your GitHub fork
2. **The `work/` directory**: Separates your work from course materials that get updated

**Files in `source/` may be:**
- Overwritten during updates
- Modified to fix issues or add clarifications  
- Removed if lessons are restructured

**Best practices:**
- Always work in the `work/` directory
- Commit and push your work regularly to your fork
- If you accidentally work in `source/`, copy changes to `work/` before pulling updates

## 📞 Getting Help

If you encounter any issues with:
- Repository structure
- Missing or corrupted files
- Git-related problems
- Exercise instructions

Please reach out during office hours or post on the course discussion forum.

## 🎯 Best Practices

1. **Collaborate and learn**: Pair program with teammates, share approaches, and learn from each other
2. **Maintain your own fork**: Each student should have their own repository for progress tracking
3. **Sync frequently**: Pull updates from upstream at the beginning of each class
4. **Copy before you code**: Always copy exercises to your work directory first  
5. **Commit regularly**: Push your work to your fork frequently to avoid losing progress
6. **Organize your work**: Keep your `work/lessons/` directory organized and well-documented
7. **Use meaningful commit messages**: Help yourself and instructors understand your progress
8. **Share solutions**: Compare different approaches with teammates after completing exercises

### Example Daily Workflow:
```bash
# Start of class - get latest materials
git fetch upstream
git merge upstream/main
git push origin main

# Work on exercises in work/ directory
cp -r source/lessons/lesson-X/ work/lessons/lesson-X/
cd work/lessons/lesson-X/
# ... do your work ...

# Save your progress
git add work/lessons/lesson-X/
git commit -m "Complete lesson-X exercise on [topic]"
git push origin main
```

---

**Happy coding, and welcome to CS 5500!** 🚀

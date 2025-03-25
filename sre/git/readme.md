# **Git: A Complete Guide from Basic to Advanced**  

Git is a **version control system (VCS)** that helps developers track changes in code, collaborate efficiently, and manage different versions of a project. It is **distributed**, meaning every developer has a full copy of the repository.

---

## **1. What is Git?**
Git is an open-source **distributed version control system (DVCS)** created by **Linus Torvalds** in 2005 for Linux kernel development.

### **Why Use Git?**
✅ Tracks changes in code over time  
✅ Enables team collaboration  
✅ Supports branching for parallel development  
✅ Works offline and is **distributed**  
✅ Integrates with platforms like GitHub, GitLab, Bitbucket  

---

## **2. Installing Git**
### **Windows**
1. Download from [git-scm.com](https://git-scm.com/downloads)
2. Install and verify:
   ```sh
   git --version
   ```

### **Mac (Using Homebrew)**
```sh
brew install git
```

### **Linux (Debian/Ubuntu)**
```sh
sudo apt update
sudo apt install git
```

---

## **3. Git Configuration**
Set up Git with your name and email:
```sh
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"
```
Verify configuration:
```sh
git config --list
```

---

## **4. Basic Git Commands**
### **Initializing a Repository**
```sh
git init
```
- Creates a new `.git` folder in the project.

### **Cloning an Existing Repository**
```sh
git clone <repository_url>
```
Example:
```sh
git clone https://github.com/user/repo.git
```

### **Checking Repository Status**
```sh
git status
```
- Shows modified, staged, and untracked files.

### **Adding Files to Staging**
```sh
git add <file>
git add .   # Add all files
```

### **Committing Changes**
```sh
git commit -m "Your commit message"
```
- Saves the changes with a message.

### **Viewing Commit History**
```sh
git log
```
- Shows commit history with IDs.

### **Checking Differences**
```sh
git diff
```
- Shows changes in modified files.

---

## **5. Working with Branches**
### **Creating a New Branch**
```sh
git branch <branch_name>
```
Example:
```sh
git branch feature-login
```

### **Switching Branches**
```sh
git checkout <branch_name>
```
or (Newer version)
```sh
git switch <branch_name>
```

### **Creating and Switching to a Branch**
```sh
git checkout -b <branch_name>
```
or
```sh
git switch -c <branch_name>
```

### **Merging Branches**
```sh
git checkout main
git merge <branch_name>
```

### **Deleting a Branch**
```sh
git branch -d <branch_name>
```
- Use `-D` for force deletion.

---

## **6. Remote Repositories**
### **Adding a Remote Repository**
```sh
git remote add origin <repository_url>
```
Example:
```sh
git remote add origin https://github.com/user/repo.git
```

### **Viewing Remote Repositories**
```sh
git remote -v
```

### **Pushing Code to a Remote Repository**
```sh
git push -u origin <branch_name>
```
Example:
```sh
git push -u origin main
```

### **Pulling Latest Changes**
```sh
git pull origin <branch_name>
```

### **Fetching Remote Changes**
```sh
git fetch
```
- Fetches changes but does not merge them.

---

## **7. Handling Merge Conflicts**
When two developers modify the same file, a conflict may occur.

### **Steps to Resolve a Conflict**
1. Run `git status` to check conflicting files.
2. Open the file and manually resolve the conflict.
3. Mark the conflict as resolved:
   ```sh
   git add <file>
   ```
4. Commit the changes:
   ```sh
   git commit -m "Resolved merge conflict"
   ```

---

## **8. Undoing Changes**
### **Discard Unstaged Changes**
```sh
git checkout -- <file>
```
or (Newer version)
```sh
git restore <file>
```

### **Unstage a File**
```sh
git reset HEAD <file>
```

### **Undo the Last Commit**
```sh
git reset --soft HEAD~1   # Keeps changes in working directory
git reset --hard HEAD~1   # Deletes changes permanently
```

---

## **9. Git Stash (Temporary Storage)**
### **Save Uncommitted Changes**
```sh
git stash
```

### **List Stashes**
```sh
git stash list
```

### **Apply Stash**
```sh
git stash apply
```

### **Delete Stash**
```sh
git stash drop
```

---

## **10. Advanced Git Commands**
### **Rebasing (Rewriting Commit History)**
```sh
git rebase <branch_name>
```
- Moves commits from one branch to another.

### **Cherry Picking (Applying a Specific Commit)**
```sh
git cherry-pick <commit_id>
```

### **Squashing Commits (Combine Multiple Commits)**
```sh
git rebase -i HEAD~3
```
- Allows merging last 3 commits.

### **Git Hooks (Automate Tasks)**
Hooks allow running scripts before/after Git actions.  
Example: Pre-commit hook (`.git/hooks/pre-commit`):
```sh
#!/bin/sh
echo "Running pre-commit hook..."
```

---

## **11. Git Tags**
### **Creating a Tag**
```sh
git tag v1.0.0
```
### **Pushing a Tag**
```sh
git push origin v1.0.0
```

### **Listing Tags**
```sh
git tag
```

### **Deleting a Tag**
```sh
git tag -d v1.0.0
```

---

## **12. Git Ignore File (`.gitignore`)**
Exclude files from Git tracking.  
Example `.gitignore`:
```
# Ignore logs and environment files
*.log
.env
node_modules/
```

---

## **13. Git Best Practices**
✅ **Write Clear Commit Messages**  
✅ **Use Feature Branches**  
✅ **Pull Before Pushing**  
✅ **Use `.gitignore` to Avoid Unwanted Files**  
✅ **Rebase Instead of Merge When Needed**  

---

## **14. Git with GitHub**
### **Step 1: Create a New Repository on GitHub**
1. Go to [GitHub](https://github.com/)
2. Click "New Repository"
3. Copy the repository URL

### **Step 2: Push Local Project to GitHub**
```sh
git remote add origin <repository_url>
git branch -M main
git push -u origin main
```

### **Step 3: Fork & Pull Request**
1. **Fork a Repository**  
   - Click the **Fork** button on GitHub.
2. **Clone the Forked Repo**  
   ```sh
   git clone <your_forked_repo_url>
   ```
3. **Create a Feature Branch & Make Changes**
   ```sh
   git checkout -b new-feature
   ```
4. **Push Changes to Your Fork**
   ```sh
   git push origin new-feature
   ```
5. **Create a Pull Request (PR) on GitHub**

---

## **15. Git vs Other Version Control Systems**
| Feature        | Git  | SVN (Subversion) | Mercurial |
|---------------|------|-----------------|-----------|
| **Type**      | Distributed | Centralized | Distributed |
| **Speed**     | Fast | Slow | Medium |
| **Branching** | Easy & Lightweight | Heavy & Costly | Moderate |
| **Offline Work** | Yes | No | Yes |

---

## **Conclusion**
Git is an **essential tool** for developers. By mastering Git, you can efficiently **manage code, collaborate, and contribute to open-source projects**.


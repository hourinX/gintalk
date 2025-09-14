#!/usr/bin/env bash
# fix-git-remote.sh
# 功能：自动检测并替换远程仓库地址（旧仓库名 → 新仓库名）
# 支持 Linux / Mac / Windows (Git Bash / WSL)

# ---------- 配置 ----------
OLD_NAME="gin-online-chat"   # 旧仓库名
NEW_NAME="gintalk"           # 新仓库名
USER="hourinX"               # GitHub 用户名
USE_SSH=true                 # true = 使用 SSH, false = 使用 HTTPS
# -------------------------

# 获取当前仓库目录
if ! git rev-parse --git-dir > /dev/null 2>&1; then
  echo "❌ 当前目录不是 Git 仓库"
  exit 1
fi

# 获取当前 origin 地址
REMOTE_URL=$(git remote get-url origin)
if [ $? -ne 0 ]; then
  echo "❌ 没有找到远程仓库 origin"
  exit 1
fi

echo "当前远程地址: $REMOTE_URL"

# 检查是否包含旧仓库名
if [[ "$REMOTE_URL" == *"$OLD_NAME"* ]]; then
  if [ "$USE_SSH" = true ]; then
    NEW_URL="git@github.com:$USER/$NEW_NAME.git"
  else
    NEW_URL="https://github.com/$USER/$NEW_NAME.git"
  fi
  git remote set-url origin "$NEW_URL"
  echo "✅ 已替换远程仓库地址为: $NEW_URL"
else
  echo "ℹ️ 远程地址中未发现 $OLD_NAME，无需修改"
fi

# 显示最终远程地址
git remote -v

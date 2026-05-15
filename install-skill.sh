#!/bin/sh
set -e

SKILL_DIR="${HOME}/.claude/skills/canvas-cli"
SKILL_URL="https://raw.githubusercontent.com/simple-scalable-solutions/canvas-cli/main/SKILL.md"

mkdir -p "$SKILL_DIR"
curl -sSfL "$SKILL_URL" -o "${SKILL_DIR}/SKILL.md"

echo "canvas-cli skill installed to ${SKILL_DIR}"
echo "Restart Claude Code to activate it."

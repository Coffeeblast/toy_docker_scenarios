SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
${SCRIPT_DIR}/.venv/bin/fastapi dev ${SCRIPT_DIR}/main.py
curl http://localhost:8000/add-story \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "9","author": "Some Dude"}'
echo ''
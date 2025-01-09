curl http://localhost:8000/delete-stories \
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE" \
    --data '{"ids" : [0, 5]}'
echo ''
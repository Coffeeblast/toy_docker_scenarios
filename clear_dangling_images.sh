docker rmi $(docker image list -f "dangling=true" -q)
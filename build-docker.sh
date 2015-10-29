#/var/bin/sh
DOCKERMACHINE_STATUS=$(docker-machine ls | grep "default")

trim() {
    local var="$*"
    var="${var#"${var%%[![:space:]]*}"}"   # remove leading whitespace characters
    var="${var%"${var##*[![:space:]]}"}"   # remove trailing whitespace characters
    echo "$var"
}

 eval "echo ~$USER"
  echo "Creating docker machine 'default'"
  $(docker-machine create --driver generic --generic-ip-address localhost "default")
  echo "Docker machine connection 'default' created"


echo "Setting Docker machine env!"
eval "$(docker-machine env default)"
docker-compose up

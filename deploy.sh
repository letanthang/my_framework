APP_NAME=$(basename $PWD)
export APP_NAME=$(echo $APP_NAME|sed -e 's/_/-/g'|cut -c1-7)
echo $APP_NAME
function build() {
  # registry_login
  
  echo "Building Dockerfile-based application..."
  # build image
  docker build \
    --build-arg HTTP_PROXY="$HTTP_PROXY" \
    --build-arg http_proxy="$http_proxy" \
    --build-arg HTTPS_PROXY="$HTTPS_PROXY" \
    --build-arg https_proxy="$https_proxy" \
    --build-arg FTP_PROXY="$FTP_PROXY" \
    --build-arg ftp_proxy="$ftp_proxy" \
    --build-arg NO_PROXY="$NO_PROXY" \
    --build-arg no_proxy="$no_proxy" \
    -t "letanthang/$APP_NAME" .
  
  docker login -u letanthang --password $DOCKER_REGISTRY_KEY
  docker push letanthang/$APP_NAME

  cd provisioning/
  sed 's/APP_NAME/'"$APP_NAME"'/g' k8s/* > zdeployment.yaml
  cat zdeployment.yaml
  kubectl apply -f zdeployment.yaml
}
build
#/bin/sh

IMAGE_NAME=$1
if [ -z "$IMAGE_NAME" ]; then
  echo "Usage: $0 <image_name>"
  exit 1
fi

# echo "https://static.wixstatic.com/media/$IMAGE_NAME/v1/fill/w_339,h_191,q_90/$IMAGE_NAME"

echo "https://static.wixstatic.com/media/$IMAGE_NAME/v1/fill/w_918,h_518,q_90/$IMAGE_NAME"
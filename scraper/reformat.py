import json, sys

def reformat_items(input_json):
    # Assuming 'input_json' is a dictionary loaded from JSON similar to the provided input
    apps_warmup_data = input_json.get("appsWarmupData", {})
    reformatted_items = []

    # Loop through each app data to find gallery data
    for app_data in apps_warmup_data.values():
        gallery_data = app_data.get("comp-kqxtj197_galleryData", {}).get("items", [])

        # Loop through each item and reformat it
        for item in gallery_data:
            metaData = item.get("metaData", {})
            reformatted_item = {
                "name": metaData.get("name", ""),
                "videoUrl": metaData.get("videoUrl", ""),
                "videoId": metaData.get("videoId", ""),
                "customPoster": metaData.get("customPoster", {}).get("url", ""),
                "mediaUrl": item.get("mediaUrl", "")
            }
            reformatted_items.append(reformatted_item)
    
    return reformatted_items

# Example usage
if __name__ == "__main__":
    input_json = json.load(sys.stdin)
    reformatted = reformat_items(input_json)
    print(json.dumps(reformatted, indent=4))

#!/bin/bash

# Directory containing the .proto files
PROTO_DIR="/proto/messages"

# Loop over all .proto files in the directory and upload them
for proto_file in $PROTO_DIR/*.proto; do
  if [ -f "$proto_file" ]; then
    # Extract the subject name from the file name (e.g., user.proto -> user-value)
    subject_name=$(basename "$proto_file" .proto)
    
    echo "Registering schema for subject: $subject_name from $proto_file"

    # Read the .proto file contents
    proto_content=$(cat "$proto_file")
    
    # Create the JSON payload for the schema registration
    json_payload=$(jq -n \
      --arg schema "$proto_content" \
      '{schema: $schema, schemaType:"PROTOBUF"}')

    # Register the .proto schema using the Schema Registry API
    curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" \
      -d "$json_payload" \
      http://schema-registry:8081/subjects/$subject_name-value/versions
  fi
done

echo "All schemas uploaded."

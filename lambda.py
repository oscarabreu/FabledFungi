import boto3
import json
import random

s3 = boto3.client('s3')

def lambda_handler(event, context):
    bucket_name = 'mushrooms3'  # Your S3 bucket name
    file_key = 'mushrooms/'  # Key of the current mushroom JSON file

    try:
        # List the JSON files in the bucket
        response = s3.list_objects_v2(Bucket=bucket_name, Prefix=file_key)
        files = response.get('Contents', [])

        if not files:
            return {'statusCode': 404, 'body': 'No JSON files found'}

        # Select a random JSON file
        selected_file = random.choice(files)['Key']
        
        # Fetch the content of the selected JSON file
        selected_file_data = s3.get_object(Bucket=bucket_name, Key=selected_file)
        selected_file_content = selected_file_data['Body'].read().decode('utf-8')
        return {
            'statusCode': 200,
            'body': selected_file_content,
            'headers': {'Content-Type': 'application/json'}
        }

    except Exception as e:
        print(e)
        return {'statusCode': 500, 'body': 'Internal Server Error'}
        
        
   

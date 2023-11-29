import json
import boto3
import redis
import os
from botocore.exceptions import ClientError

dynamodb = boto3.resource('dynamodb', region_name='us-east-1')
s3_client = boto3.client('s3')

REDIS_HOST = '172.31.20.70'
REDIS_PORT = 6379
TABLE_NAME = 'Mushroom'

def generate_presigned_url(s3_uri, expiration=3600):
    """Generate a presigned URL for an S3 object."""
    try:
        bucket, key = s3_uri.replace("s3://", "").split("/", 1)
        
        url = s3_client.generate_presigned_url('get_object',
                                               Params={'Bucket': bucket, 'Key': key},
                                               ExpiresIn=expiration)
        return url
    except ClientError as e:
        print(f"Error generating presigned URL: {e}")
        return None
    
def lambda_handler(event, context):
    r = redis.Redis(host=REDIS_HOST, port=REDIS_PORT, decode_responses=True)

    random_key = r.randomkey()
    json_data = r.get(random_key)

    if json_data:
        data = json.loads(json_data)
        source_species = data.get('primaryKey')
        observation = data.get('sortKey')
        
        if source_species and observation:
            table = dynamodb.Table(TABLE_NAME)
            try:
                response = table.get_item(Key={'Source#Species': source_species, 'Observation': observation})
                item = response.get('Item')

                if item:
                    if 'ImageURL' in item:
                        presigned_url = generate_presigned_url(item['ImageURL'])
                        if presigned_url:
                            item['ImageURL'] = presigned_url
                        else:
                            return {'statusCode': 500, 'body': 'Error generating presigned URL'}

                    return {
                        'statusCode': 200,
                        'body': json.dumps(item)
                    }
                else:
                    return {
                        'statusCode': 404,
                        'body': 'Item not found in DynamoDB'
                    }
            except Exception as e:
                return {
                    'statusCode': 500,
                    'body': f"Error fetching item from DynamoDB: {str(e)}"
                }
        else:
            return {
                'statusCode': 500,
                'body': "primaryKey or sortKey key is missing in the Redis data"
            }
    else:
        return {
            'statusCode': 404,
            'body': 'Random key not found in Redis'
        }

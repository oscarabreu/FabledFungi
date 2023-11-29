# FabledFungi v2

## Summary

FabledFungi is a unique API hosted on AWS, born from a rich dataset of mushroom observations, spanning from 2008 to 2015, sourced from iNaturalist. This dataset was originally in CSV format and has been significantly enhanced with custom scripts. These scripts are designed to extract and augment various details such as observation dates, locations, names, and associated image URLs, enriching the original data.

At its core, FabledFungi is designed to fetch random mushroom observation data, providing users not just with the image URL but also with a wealth of related metadata for each observation.

This project represents the culmination of two primary objectives. Firstly, it served as a practical application of the knowledge I acquired while studying for the AWS Solutions Architect certification. Secondly, it offered a hands-on experience in backend web development. I've made all resources, including data, scripts, and web development tools, openly available for anyone interested in exploring, learning, or building upon this project.

**Concepts**: AWS Virtual Private Cloud, CORS, VPC Endpoints, Security Groups, NAT Gateway, API Gateway, S3, pre-Signed URL, Lambda (Python 11), Eventbridge, DynamoDB On-Demand, UURI, EC2 (Amazon Linux), Redis, Docker, Go Scripting, AWS Amplify, API Management, LCP, FCP, CLS, Amazon X-Ray
  
### Sample API Response:

  ```
  {"statusCode": 200, "body": "{\"UpdatedAt\": \"2021-01-25 22:34:45 UTC\", \"TaxonFamilyName\": \"Amanitaceae\", \"Longitude\": \"-69.0051\", \"TaxonClassName\": \"Agaricomycetes\", \"ScientificName\":
  \"Amanita muscaria guessowii\", \"Latitude\": \"45.1595\", \"License\": \"CC-BY-NC\", \"Observation\": \"5701181\", \"TaxonKingdomName\": \"Fungi\", \"TagList\": \"\", \"ImageURL\": \"https://fungis3.s3.
  amazonaws.com/iNaturalist/5701181.jpg?AWSAccessKeyId=ASIAZRBRGY6KDBYKFOXH&Signature=F7I5vbj36%2B11Y6jjLMscINwEzK0%3D&x-amz-security-token=IQoJb3JpZ2luX2VjEIf%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FwEaCXVzLWVhc3QtMS
  JHMEUCIE%2B2ycR%2FCUk1lqmguUy0DHNnOBRvQl59deMvveb44C9%2BAiEAufyfx53b0X2WpW%2FqEZ%2FOAgiQRcEgG72B%2BVl56QzH4ZAq%2BwII3%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARAAGgw2NTUwODU3MTczOTYiDGavTfvBevo14Rsx%2FSrPAqUNuNj
  ulImJrInkJTED%2BWoZ2ftPRfZqHhdUrT6SjjPUTL%2Bndw%2FHtu1c%HELLO WORLD%2Bx3Q%2FXmTowPhTxfG0scGvt4X7nWOAE3qJb%2FioxKmhypjFxS8i02UV2sXos4XyEAGE106743%2F1HV%2F9pFOjpL7fclyKyj
  HELLO WORLD%2BzPouMhJFzqcSlFjvM5YCAA11sLc2Hhu%2FA4%HELLO WORLD%2BKfaB93tqGSH5xG25zfzeTg7QmjlgwGWBzAKwkknLvj%2FMqx1E%2BhqYCyXE5NfncGABM5zHY
  gEhDAsHv3tLuJaX2tiPoVwTgrvhApq4DUr87C6r0kBEOJHJm%2BVG0YzJw0E6cdQdVBMI3snqsGOp4BQuLDCG4%2BTYGMlA9Pdi%2BQLzwFZtlvHIGEU5DAlQs4I6fNLiTwOrIQu5GT%2FTK9gWrlbMFmhn5VZZ%2Bj5fCRi4O%2Bbqqq2CsO%2BS7U4ShO9DZ8ZBZyDale
  QJ0thA5HECrFrP4AutuKIDho86WfGUta%2BwdQNjkR7pDqhQNBIU9dwAubkr8RTW0d7uzLCzAF17lb4GV1sV2FLXYCKIz6M%2BO1nF0%3D&Expires=1701299231\", \"Source#Species\": \"iNaturalist#Amanita muscaria\", \"TaxonPhylumName\":
   \"Basidiomycota\", \"ObservedOn\": \"9/30/10\", \"CreatedAt\": \"2017-04-14 22:17:13 UTC\", \"TaxonID\": \"321524\", \"Description\": \"\", \"TaxonOrderName\": \"Agaricales\", \"URL\": \"https://www.ina
  turalist.org/observations/5701181\", \"UserID\": \"382262\", \"TaxonGenusName\": \"Amanita\", \"UserLogin\": \"erlonbailey\", \"CommonName\": \"American Yellow Fly Agaric\", \"PlaceGuess\": \"Maine, USA\
  ", \"SpeciesGuess\": \"American Yellow Fly Agaric\", \"TaxonSpeciesName\": \"Amanita muscaria\"}"}
  ```

## Architecture:
<p align="center">
  <img src="https://github.com/oscarabreu/FabledFungi/assets/99779654/4f90dfe9-a17e-4870-bf4f-750833f031d9" alt="Screenshot 2023-11-27 at 3 15 57 PM">
</p>

## Decisions:
- **Website Hosting** - This static site is hosting using AWS Amplify, a fully-managed AWS service.
- **API Taffic/Routing** - API Routing/Management is handled by the AWS API Gateway - a fully-managed AWS service.
- **API Functionality** The FabledFungiAPI is designed solely for retrieving information, which means it only supports GET requests. Cross-Origin Resource Sharing (CORS) is enabled to allow the API to be accessed from different domains.
- **Data Storage** - The metadata for all mushroom observations is maintained in a DynamoDB table. One of the columns in this table includes the S3 URI where the image is located.
- **Redis(EC2)** - For O(1) random index time, Redis was implemented to hold a UUID as the key, and the DynamoDB partitionkey/sortkey as the value.
- **DynamoDB** - To optimize query efficiency, DynamoDB is utilized for managing the metadata of mushroom observations, encompassing approximately 25 columns. In this DynamoDB table, 'Source#Species' acts as the partition key, and 'ObservationID' functions as the sort key. These keys facilitate indexing and constructing the full JSON response.
- **S3** - The S3 holds all mushoom observation images in the form of jpg. They are named by /ObservationID.jpg/, where ObservationID is the ID given by its source. 
- **Lambda** - This lambda, powered by Python 11 and boto3, peforms 3 functions:
  1. It accesses EC2 Linux (Redis) with proper Security Group, VPC, and IAM permissions to call randomkey(), and get a random DynamoDB Partition/Sort as the value.
  2. This time, with proper VPC Endpoint configs + IAM permissions, it securely accesses the DynamoDB with this random Partition/Sort key and assembles a JSON response with the >25 columns of data.
  3. Before shipping the data back through the API Gateway, it generates a pre-signed URL from the S3 ImageURL so the client (AWS Amplify / Javascript) can securely fetch the S3 image data once received.

## Improvements: 
- Largest Contentful Paint (LCP) Average ... v1 = 3.475, v2 = 0.850
- First Contentful Paint (FCP) Average ... v1 = .875, v2 = .525
- Cumulative Layout Shift (CLS) Average ...  v1 = .221, v2 = .060
- PageSpeed Performance Average ... v1 = 82, v2 = 98.5
- **Lambda Duration Average** ... v1 = 2038 ms, v2 = 225 ms
- **Image Latency Average** ... v1 = ~400ms, v2 = 310.8ms 

## Ideas for further improvements:
- Set an EventBridge to cache 100 random keys from EC2 Linux (Redis) every hour, and let the browser generate a random token from 1-100 on every get requst. Implement caching on API Gateway so users can benefit from CDN caching. Be sure to implement cache invalidation so that users are not stuck receiving the same data.
- Implement CloudFront on S3 to further reduce the latency for the image request
- "Warm" lambdas using EventBridge invocations

# Images

## Lambda Performance
![Screenshot 2023-11-29 at 5 11 57 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/1604bc34-81f2-4764-9b42-00463a3729a7)


## Frontend Performance
![Screenshot 2023-11-29 at 12 25 28 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/4f0972c6-536e-45ea-9b12-0bc84d4bb4a7)

## Network 
![Screenshot 2023-11-29 at 5 15 14 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/ba898a45-e4b8-454d-8e77-4e819be3099c)

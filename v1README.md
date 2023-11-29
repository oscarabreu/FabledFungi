# FabledFungi v1

## Summary

FabledFungi is a unique API hosted on AWS, born from a rich dataset of mushroom observations, spanning from 2008 to 2015, sourced from iNaturalist. This dataset was originally in CSV format and has been significantly enhanced with custom scripts. These scripts are designed to extract and augment various details such as observation dates, locations, names, and associated image URLs, enriching the original data.

At its core, FabledFungi is designed to fetch random mushroom observation data, providing users not just with the image URL but also with a wealth of related metadata for each observation.

This project represents the culmination of two primary objectives. Firstly, it served as a practical application of the knowledge I acquired while studying for the AWS Solutions Architect certification. Secondly, it offered a hands-on experience in backend web development. I've made all resources, including data, scripts, and web development tools, openly available for anyone interested in exploring, learning, or building upon this project.

## V1 - FabledFungis API

### Architecture:
<p align="center">
  <img src="https://github.com/oscarabreu/FabledFungi/assets/99779654/626f48ef-4891-4dcc-a5f8-8b7d40daaca9" alt="Screenshot 2023-11-27 at 3 15 57 PM">
</p>

### Decisions:
- **Website Hosting** - This static site is hosting using AWS Amplify, a fully-managed AWS service.
- **API Taffic/Routing** - API Routing/Management is handled by the AWS API Gateway - a fully-managed AWS service.
- **API Functionality** The FabledFungiAPI is designed solely for retrieving information, which means it only supports GET requests. Cross-Origin Resource Sharing (CORS) is enabled to allow the API to be accessed from different domains
- **Dataset** - All mushroom observation URL+Metadata is stored as a JSON object in an s3 bucket named fungis3.
- **Lambda** - Implemented in Python with the boto3 library, this Lambda calls ListObjectsV2, and retrieves a random index in an S3 bucket containing Mushroom Observations in the form of JSON. 

### Sample API Response:
  ```
  {"statusCode": 200, "body": "{\"taxon_id\":\"179750\",\"observer_id\":\"100940464\",\"observed_on\":\"8/30/07\",\"user_id\":\"4391274\",\"user_login\":\"cefreebury\",\"created_at\":\"2021-11-12 18:09:07
UTC\",\"updated_at\":\"2022-04-11 13:38:22 UTC\",\"license\":\"CC-BY-NC\",\"url\":\"https://www.inaturalist.org/observations/100940464\",\"image_url\":\"https://inaturalist-
open-data.s3.amazonaws.com/photos/168523427/medium.jpg\",\"tag_list\":\"Gatineau Park\",\"description\":\"On thin soil in abandoned parking lot at entrance to Trail
56.\",\"place_guess\":\"Gatineau Park, Quebec, Canada\",\"latitude\":\"45.5996314\",\"longitude\":\"-76.09688953\",\"species_guess\":\"Split-peg
Lichen\",\"scientific_name\":\"Cladonia cariosa\",\"common_name\":\"Split-peg
Lichen\",\"taxon_kingdom_name\":\"Fungi\",\"taxon_phylum_name\":\"Ascomycota\",\"taxon_class_name\":\"Lecanoromycetes\",\"taxon_order_name\":\"Lecanorales\",\"taxon_family_name\
":\"Cladoniaceae\",\"taxon_genus_name\":\"Cladonia\",\"taxon_species_name\":\"Cladonia cariosa\"}", "headers": {"Content-Type": "application/json"}}
  ```

### Lambda Performance

![Screenshot 2023-11-26 at 6 27 55 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/c3e7fa36-694c-4769-b073-0aa6ed4372cc)

### Frontend Performance
![FabledFungi v1 - SS4](https://github.com/oscarabreu/FabledFungi/assets/99779654/93cee26c-7657-46b7-bc3b-09f74070e17d)

### Network 
![Screenshot 2023-11-27 at 1 33 50 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/5481c3dc-3e52-4992-a06a-18648813b6e7)


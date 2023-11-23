# FabledFungi
## API Deployment

v1 of API managed by AWS Eventbridge, API Gateway (GET/OPTION CORS Enabled), two AWS Lambdas - one invoked by Eventbridge another invoked by Gateway.

v2 of API managed by API Gateway (GET/OPTION CORS Enabled), one AWS Lambda invoked by Gateway. 

## Function:

S3 Bucket holds iNaturalist.nz Mushroom Observation data from 2008 - 2015. 

Lambda/API outputs a random "Mushroom" from this S3 in the form of JSON data after each invocation.

Content is displayed in FabledFungis.com [I know _fungi_ is the plural term. Please don't yell at me.] 

 ## Example:
  ```
  {"statusCode": 200, "body": "{\"taxon_id\":\"179750\",\"observer_id\":\"100940464\",\"observed_on\":\"8/30/07\",\"user_id\":\"4391274\",\"user_login\":\"cefreebury\",\"created_at\":\"2021-11-12 18:09:07
UTC\",\"updated_at\":\"2022-04-11 13:38:22 UTC\",\"license\":\"CC-BY-NC\",\"url\":\"https://www.inaturalist.org/observations/100940464\",\"image_url\":\"https://inaturalist-
open-data.s3.amazonaws.com/photos/168523427/medium.jpg\",\"tag_list\":\"Gatineau Park\",\"description\":\"On thin soil in abandoned parking lot at entrance to Trail
56.\",\"place_guess\":\"Gatineau Park, Quebec, Canada\",\"latitude\":\"45.5996314\",\"longitude\":\"-76.09688953\",\"species_guess\":\"Split-peg
Lichen\",\"scientific_name\":\"Cladonia cariosa\",\"common_name\":\"Split-peg
Lichen\",\"taxon_kingdom_name\":\"Fungi\",\"taxon_phylum_name\":\"Ascomycota\",\"taxon_class_name\":\"Lecanoromycetes\",\"taxon_order_name\":\"Lecanorales\",\"taxon_family_name\
":\"Cladoniaceae\",\"taxon_genus_name\":\"Cladonia\",\"taxon_species_name\":\"Cladonia cariosa\"}", "headers": {"Content-Type": "application/json"}}
  ```

My goal is to optimize this process as much as possible.

Template Frontend:

![Screenshot 2023-11-23 at 5 29 59 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/4262dca9-559f-4678-bfb5-432207cad2bc)
![Screenshot 2023-11-23 at 5 30 06 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/3a452a19-189d-4606-8f7c-37ffd1e01d0f)
![Screenshot 2023-11-23 at 5 30 19 PM](https://github.com/oscarabreu/FabledFungi/assets/99779654/ad047f36-af46-44e2-b19b-1849444852b2)

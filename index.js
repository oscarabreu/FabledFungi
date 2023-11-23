document.addEventListener("DOMContentLoaded", function () {
    // Your code here
    const url = "https://8n31nzke69.execute-api.us-east-1.amazonaws.com/v1";
    
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        const apiResponse = JSON.parse(data.body);
        console.log(apiResponse); // Log the JSON data to the console
        const imageElement = document.querySelector("img");
        imageElement.src = apiResponse.image_url;

    // Select HTML elements by class name and fill them with data
    document.querySelector(".Mushroom-Name").textContent = apiResponse.taxon_species_name;
    document.querySelector(".Observed b").textContent = apiResponse.observed_on;
    document.querySelector(".Location b").textContent = apiResponse.place_guess;
    document.querySelector(".Latitude b").textContent = apiResponse.latitude;
    document.querySelector(".Longitude b").textContent = apiResponse.longitude;
    document.querySelector(".Kingdom-Name b").textContent = apiResponse.taxon_kingdom_name;
    document.querySelector(".Phylum-Name b").textContent = apiResponse.taxon_phylum_name;
    document.querySelector(".Class-Name b").textContent = apiResponse.taxon_class_name;
    document.querySelector(".Order-Name b").textContent = apiResponse.taxon_order_name;
    document.querySelector(".Family-Name b").textContent = apiResponse.taxon_family_name;
    document.querySelector(".Genus-Name b").textContent = apiResponse.taxon_genus_name;
    document.querySelector(".Species-Name b").textContent = apiResponse.taxon_species_name;
    const observer_id = apiResponse.observer_id.toString(); // Convert the ID to a string
    document.querySelector(".source-link").href = `https://www.inaturalist.org/observations/${observer_id}`; // Set the href attribute
    document.querySelector(".User-Login b").textContent = apiResponse.user_login;
    document.querySelector(".Description b").textContent = apiResponse.description;

})
  .catch((error) => {
    console.error('Error loading data:', error); // Log an error message to the console
  });
});
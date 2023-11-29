document.addEventListener("DOMContentLoaded", function () {
    const url = "https://rghlagldqe.execute-api.us-east-1.amazonaws.com/v2";
    
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        const apiResponse = JSON.parse(data.body);
        console.log(apiResponse); 
        loadImageAsync(apiResponse.ImageURL);
        document.querySelector(".Mushroom-Name").textContent = apiResponse.TaxonSpeciesName;
        document.querySelector(".Observed b").textContent = apiResponse.ObservedOn;
        document.querySelector(".Location b").textContent = apiResponse.PlaceGuess;
        document.querySelector(".Latitude b").textContent = apiResponse.Latitude;
        document.querySelector(".Longitude b").textContent = apiResponse.Longitude;
        document.querySelector(".Kingdom-Name b").textContent = apiResponse.TaxonKingdomName;
        document.querySelector(".Phylum-Name b").textContent = apiResponse.TaxonPhylumName;
        document.querySelector(".Class-Name b").textContent = apiResponse.TaxonClassName;
        document.querySelector(".Order-Name b").textContent = apiResponse.TaxonOrderName;
        document.querySelector(".Family-Name b").textContent = apiResponse.TaxonFamilyName;
        document.querySelector(".Genus-Name b").textContent = apiResponse.TaxonGenusName;
        document.querySelector(".Species-Name b").textContent = apiResponse.TaxonSpeciesName;
        const observerid = apiResponse.observerid.toString(); // Convert the ID to a string
        document.querySelector(".source-link").href = `https://www.inaturalist.org/observations/${observerid}`; // Set the href attribute
        document.querySelector(".User-Login b").textContent = apiResponse.userlogin;
        document.querySelector(".Description b").textContent = apiResponse.description;

})
  .catch((error) => {
    console.error('Error loading data:', error); 
  });
});


function loadImageAsync(url) {
    const imageElement = document.querySelector("img");
    const tempImage = new Image();

    tempImage.onload = function () {
        imageElement.src = this.src;
    };

    tempImage.src = url;
}
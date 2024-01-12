<script setup>

</script>

<template>
  <div class="row justify-content-center">
    <div class="col-12 col-lg-10 d-flex flex-column justify-content-center gap-5">
      <div id="carouselExampleIndicators" class="carousel slide" data-bs-ride="true">
        <div class="carousel-indicators">
          <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="0" class="active"
            aria-current="true" aria-label="Slide 1"></button>
          <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="1"
            aria-label="Slide 2"></button>
          <button type="button" data-bs-target="#carouselExampleIndicators" data-bs-slide-to="2"
            aria-label="Slide 3"></button>
        </div>
        <div class="carousel-inner">
          <div class="carousel-item active">
            <img src="../assets/homepage/pexels-eduardo-rosas-907489.jpg" class="d-block w-100 rounded homeImg" alt="...">
          </div>
          <div class="carousel-item">
            <img src="../assets/homepage/pexels-pixabay-267350.jpg" class="d-block w-100 rounded homeImg" alt="...">
          </div>
          <div class="carousel-item">
            <img src="../assets/homepage/pexels-plann-4549414.jpg" class="d-block w-100 rounded homeImg" alt="...">
          </div>
        </div>
        <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleIndicators"
          data-bs-slide="prev">
          <span class="carousel-control-prev-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Previous</span>
        </button>
        <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleIndicators"
          data-bs-slide="next">
          <span class="carousel-control-next-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Next</span>
        </button>
      </div>
      <div class="col">
        <h3 class="text-success opacity-75">Google Play Store</h3>
        <p>
          <i>Google Play Store</i> filters the numbers of apps as per user's need. User can filter apps according to application's category and whether the app is free or paid. No doubt user can find required application by searching it. It not only offers features like adding apps to store, but also  updating it by filling required details. If user wants to delete the app from store, they can do it by delete feature.
        </p>
        <p class="mb-0">
          The details page shows more information about specific application. It also Rating and Review section where it displays reviews and rating associated with app and user can add their opinion and experience about application thorough adding ratings and reviews.
        </p>
      </div>
    </div>
  </div>
</template>

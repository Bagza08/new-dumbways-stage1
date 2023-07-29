// menggunakan Functional

// fake data -> array of object
const testimonialData = [
  {
    user: "Kratos",
    quote: "Jelek banget haha",
    image:
      "https://wallpapers.com/images/hd4/god-of-war-kratos-red-tattoos-is5vckclyv5nbivb.jpg",
    rating: 3,
  },
  {
    user: "Atreus",
    quote: "mungkin kurang bagus",
    image:
      "https://oyster.ignimgs.com/mediawiki/apis.ign.com/god-of-war-ragnarok/1/1b/Ater.jpg",
    rating: 4,
  },
  {
    user: "Dumways",
    quote: "mantapp joss gandos",
    image: "img/logo.png",
    rating: 5,
  },
  {
    user: "Baldur",
    quote: "amazinggggg!!!",
    image:
      "https://api.duniagames.co.id/api/content/upload/file/14578941721580450975.JPG",
    rating: 5,
  },
];

// let testimonialHTML = ``;

// // foreach
// testimonialData.forEach(function (card) {
//   testimonialHTML += `<div class="card m-3 shadow-sm" style="width: 18rem; object-fit: cover">
//     <img
//         src="${card.image}"
//         class="card-img-top object-fit-cover"
//         height="180rem"
//     />
//     <div class="card-body">
//         <p class="card-text text-start fst-italic">"${card.quote}
//         "</p>
//         <div class="d-flex justify-content-end">
//         <h6 class="card-title text-start">- ${card.user}</h6>
//         </div>
//         <div class="d-flex justify-content-end">
//         <h6 class="card-title text-start">${card.rating} <i class="bi bi-star-fill"></i></h6>
//         </div>
//     </div>
//     </div>`;
// });

// document.getElementById("testimonial").innerHTML = testimonialHTML;

function allTestimonial() {
  let testimonialHTML = ``;

  // foreach
  testimonialData.forEach(function (card) {
    testimonialHTML += `<div class="card m-3 shadow-sm" style="width: 18rem; object-fit: cover">
      <img
          src="${card.image}"
          class="card-img-top object-fit-cover"
          height="180rem"
      />
      <div class="card-body">
          <p class="card-text text-start fst-italic">"${card.quote}
          "</p>
          <div class="d-flex justify-content-end">
          <h6 class="card-title text-start">- ${card.user}</h6>
          </div>
          <div class="d-flex justify-content-end">
          <h6 class="card-title text-start">${card.rating} <i class="bi bi-star-fill"></i></h6>
          </div>
      </div>
      </div>`;
  });

  document.getElementById("testimonial").innerHTML = testimonialHTML;
}

// eksekusi func all
allTestimonial();

function filterTestimonial(rating) {
  let filteredTestimonialHTML = ``;

  const filteredData = testimonialData.filter((card) => {
    return card.rating === rating;
  });

  filteredData.forEach((card) => {
    filteredTestimonialHTML += `<div class="card m-3 shadow-sm" style="width: 18rem; object-fit: cover">
    <img
        src="${card.image}"
        class="card-img-top object-fit-cover"
        height="180rem"
    />
    <div class="card-body">
        <p class="card-text text-start fst-italic">"${card.quote}
        "</p>
        <div class="d-flex justify-content-end">
        <h6 class="card-title text-start">- ${card.user}</h6>
        </div>
        <div class="d-flex justify-content-end">
        <h6 class="card-title text-start">${card.rating} <i class="bi bi-star-fill"></i></h6>
        </div>
    </div>
    </div>`;
  });

  document.getElementById("testimonial").innerHTML = filteredTestimonialHTML;
}

// map
// const testimonialMap = testimonialData.map(function (card) {
//   return `<div class="card m-3" style="width: 18rem; object-fit: cover">
//       <img
//           src="${card.image}"
//           class="card-img-top object-fit-cover"
//           height="180rem"
//       />
//       <div class="card-body">
//           <p class="card-text text-start fst-italic">"${card.quote}
//           "</p>
//           <div class="d-flex justify-content-end">
//           <h6 class="card-title text-start">- ${card.user}</h6>
//           </div>
//       </div>
//       </div>`;
// });

// console.log(testimonialMap); // bentuk array
// console.log(testimonialMap.join()); // bentuk string

// document.getElementById("testimonial").innerHTML = testimonialMap.join();

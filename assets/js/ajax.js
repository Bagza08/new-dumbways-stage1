const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();

  xhr.open("GET", "https://api.npoint.io/c1804081e62c479be1c4", true);

  xhr.onload = function () {
    if (xhr.status === 200) {
      // nilai 200 adalah berhasil (cari tahu di internet menggunakan http status code)
      //   console.log(xhr.response);
      //   console.log(xhr.responseText);
      resolve(JSON.parse(xhr.response));
      //JSON.parse ----> yaitu bawaan dari js untuk parsing bentukan string menjadi sebuah JSON
    } else if (xhr.status >= 400) {
      reject("error load data");
    }
  };

  xhr.onerror = function () {
    reject("network error");
  };

  xhr.send();
});

// promise-chaining
// promise
//   .then((value) => {
//     console.log(value);
//   })
//   .catch((err) => {
//     console.log(err);
//   });

let testimonialData = [];

// async-await
async function getData() {
  try {
    const respone = await promise;
    testimonialData = respone;
    allTestimonial();
    console.log(respone);
  } catch (err) {
    console.log(err);
  }
}

getData();

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

  console.log(testimonialData);
}

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

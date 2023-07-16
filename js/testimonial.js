// ini adalah class concept ✅
class Testimonial {
  //   encapsulation ✅
  #image = "";
  #quote = "";

  constructor(quote, image) {
    // this.user = user;
    this.#quote = quote;
    this.#image = image;
  }

  // abstraction ✅
  get quote() {
    return this.#quote;
  }

  get image() {
    return this.#image;
  }

  get user() {
    // return "";
    throw new Error("there is must be user to make testimonials");
  }

  get testimonialHTML() {
    return `
    <div class="card m-3" style="width: 18rem; object-fit: cover">
      <img
          src="${this.image}"
          class="card-img-top object-fit-cover"
          height="180rem"
      />
      <div class="card-body">
          <p class="card-text text-start fst-italic">"${this.quote}
          "</p>
          <div class="d-flex justify-content-end">
          <h6 class="card-title text-start">- ${this.user}</h6>
          </div>
      </div>
      </div>
    `;
  }
}

// ini adalah inheritance ✅
class userTestimonial extends Testimonial {
  #user = "";
  constructor(user, quote, image) {
    super(quote, image);
    this.#user = user;
  }

  //   polymorphism ✅
  get user() {
    return "user :" + this.#user;
  }
}

// ini adalah inheritance ✅
class companyTestimonial extends Testimonial {
  #company = "";
  constructor(company, quote, image) {
    super(quote, image);
    this.#company = company;
  }

  //   polymorphism ✅
  get user() {
    return "company :" + this.#company;
  }
}

// ini adalah object concept ✅
const testimonial1 = new companyTestimonial(
  "dumways",
  "mantapp joss gandos",
  "img/logo.png"
);

const testimonial2 = new userTestimonial(
  "Kratos",
  "Jelek banget haha",
  "https://wallpapers.com/images/hd4/god-of-war-kratos-red-tattoos-is5vckclyv5nbivb.jpg"
);

const testimonial3 = new userTestimonial(
  "Atreus",
  "mungkin kurang bagus",
  "https://oyster.ignimgs.com/mediawiki/apis.ign.com/god-of-war-ragnarok/1/1b/Ater.jpg"
);

// const testimonial4 = new Testimonial(
//   "mungkin kurang bagus",
//   "https://oyster.ignimgs.com/mediawiki/apis.ign.com/god-of-war-ragnarok/1/1b/Ater.jpg"
// );

let testimonialData = [testimonial1, testimonial2, testimonial3];

let testimonialHTML = "";

for (let index = 0; index < testimonialData.length; index++) {
  testimonialHTML += testimonialData[index].testimonialHTML;
}

document.getElementById("testimonial").innerHTML = testimonialHTML;

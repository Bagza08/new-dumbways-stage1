let dataMyproject = [];

// function button addproject
function addProject(event) {
  event.preventDefault();

  let pName = document.getElementById("iproject").value;
  let sDate = new Date(document.getElementById("sdate").value);
  let eDate = new Date(document.getElementById("edate").value);

  // jarak input sdate dan edate dalam milisecond
  let difference = "";

  if (sDate > eDate) {
    difference = sDate - eDate;
  } else if (sDate < eDate) {
    difference = eDate - sDate;
  }
  // ke detik
  let second = Math.round(difference / 1000);
  // ke menit
  let minute = Math.round(second / 60);
  // ke jam
  let hour = Math.round(minute / 60);
  // ke hari
  let day = Math.round(hour / 24);
  // ke bulan
  let month = Math.round(day / 30);

  let duration = "";
  let durationmore = "";

  if (day < 30) {
    duration = `${day} Hari`;
  } else if (day == 30) {
    duration = `${month} Bulan`;
  } else if (day > 30) {
    durationmore = day % 30;
    duration = `${month} Bulan ${durationmore} Hari`;
  }

  console.log(sDate);
  // console.log(durationmore);
  // console.log(typeof difference);

  let desc = document.getElementById("desc").value;

  const reactjs = '<i class="fa-brands fa-react fa-2xl marlef"></i>';
  const nodejs = '<i class="fa-brands fa-node fa-2xl marlef"></i>';
  const bootstrap = '<i class="fa-brands fa-bootstrap fa-2xl marlef"></i>';
  const laravel = '<i class="fa-brands fa-laravel fa-2xl marlef"></i>';

  let reJs = document.getElementById("rejs").checked ? reactjs : "";
  let noJs = document.getElementById("nojs").checked ? nodejs : "";
  let bs = document.getElementById("bs").checked ? bootstrap : "";
  let lar = document.getElementById("lar").checked ? laravel : "";

  let image = document.getElementById("title-attfile").files;

  image = URL.createObjectURL(image[0]);

  let project = {
    pName,
    duration,
    desc,
    reJs,
    noJs,
    bs,
    lar,
    image,
  };

  dataMyproject.push(project);
  renderProject();

  console.log(dataMyproject);
}

function renderProject() {
  document.getElementById("cardcontainer").innerHTML = "";

  for (let index = 0; index < dataMyproject.length; index++) {
    document.getElementById("cardcontainer").innerHTML += `
    <div class="card m-3" style="width: 18rem; object-fit: cover">
    <img
      src="${dataMyproject[index].image}"
      class="card-img-top object-fit-cover"
      height="180rem"
    />
    <div class="card-body">
      <p class="card-text text-start h6 f-1">
      ${dataMyproject[index].pName}
      </p>
      <p class="card-text text-start h6 f-2">
        Durasi : ${dataMyproject[index].duration}
      </p>
      <p class="card-text text-start h6 f-2 mb-5 mt-4">
      ${dataMyproject[index].desc}
      </p>
      <div class="d-flex justify-content-start my-4">
      <span >${dataMyproject[index].reJs}</span>
      <span class="mx-2">${dataMyproject[index].noJs}</span>
      <span class="mx-2">${dataMyproject[index].bs}</span>
      <span class="mx-2">${dataMyproject[index].lar}</span> 
      </div>
      <div class="d-flex justify-content-evenly mt-5" >
        <a href="" class="btn btn-primary col-5" id="btn-b" style="border-radius: 20px;">Edit</a>
        <a href="" class="btn btn-primary col-5" id="btn-b" style="border-radius: 20px;">Delete</a>
      </div>
    </div>
  </div>

    `;
  }
}

// let sDate = document.getElementById("sdate");

// console.log(sDate);

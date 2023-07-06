let dataMyproject = [];

// function checkboxFix() {}

function addProject(event) {
  event.preventDefault();

  let pName = document.getElementById("iproject").value;
  //   let sDate = document.getElementById("sdate").value;
  //   let eDate = document.getElementById("edate").value;
  let desc = document.getElementById("desc").value;
  //   let reJs = document.getElementById("rejs").checked;
  //   console.log(reJs);
  //   let noJs = document.getElementById("nojs").value;
  //   let bs = document.getElementById("bs").value;
  //   let lar = document.getElementById("lar").value;

  let reJs = document.getElementById("rejs").checked;
  let noJs = document.getElementById("nojs").checked;
  let bs = document.getElementById("bs").checked;
  let lar = document.getElementById("lar").checked;

  let reactjs = "";
  let nodejs = "";
  let bootstrap = "";
  let laravel = "";

  //   if (reJs.checked == true) {
  //     reactjs = '<i class="fa-brands fa-react fa-2xl">';
  //   }

  //   if (noJs.checked == true) {
  //     nodejs = '<i class="fa-brands fa-node fa-2xl marlef">';
  //   }

  //   if (bs.checked == true) {
  //     bootstrap = '<i class="fa-brands fa-bootstrap fa-2xl marlef">';
  //   }

  //   if (lar.checked == true) {
  //     laravel = '<i class="fa-brands fa-laravel fa-2xl marlef"></i>';
  //   }

  //   console.log(reactjs);

  //   if (reJs.checked == true) {
  //     reactjs = '<i class="fa-brands fa-react fa-2xl">';
  //   } else if (noJs.checked == true) {
  //     nodejs = '<i class="fa-brands fa-node fa-2xl marlef">';
  //   } else if (bs.checked == true) {
  //     bootstrap = '<i class="fa-brands fa-bootstrap fa-2xl marlef">';
  //   } else if (lar.checked == true) {
  //     laravel = '<i class="fa-brands fa-laravel fa-2xl marlef"></i>';
  //   }

  let image = document.getElementById("title-attfile").files;

  image = URL.createObjectURL(image[0]);

  let project = {
    pName,
    // sDate,
    // eDate,
    desc,
    // reJs: reactjs,
    // noJs: nodejs,
    // bs: bootstrap,
    // lar: laravel,
    image,
  };

  dataMyproject.push(project);
  renderProject();

  //   console.log(project);
}

function renderProject() {
  document.getElementById("cardcontainer").innerHTML = "";

  for (let index = 0; index < dataMyproject.length; index++) {
    document.getElementById("cardcontainer").innerHTML += `
    <div id="card" class="card">
        <img src="${dataMyproject[index].image}" alt="" />
        <h4>${dataMyproject[index].pName}</h4>
        <h5>durasi 3 bulan</h5>
        <p>${dataMyproject[index].desc}
        </p>
        <div class="icheckbox">
            <i class="fa-brands fa-react fa-2xl"></i>
            <i class="fa-brands fa-node fa-2xl marlef"></i>
            <i class="fa-brands fa-bootstrap fa-2xl marlef"></i>
            <i class="fa-brands fa-laravel fa-2xl marlef"></i>
        </div>
        <div class="duatombol">
        <a href="">Edit</a>
        <a href="">Delete</a>
        </div>
    </div>

    `;
  }
}

// let sDate = document.getElementById("sdate");

// console.log(sDate);

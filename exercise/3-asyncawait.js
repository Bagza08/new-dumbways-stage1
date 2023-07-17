// sama seperti promise chaining (then catch) kegunaannya cuman beda penulisan

let condition = true;

let janji = new Promise((resolve, reject) => {
  if (condition) {
    setTimeout(() => {
      resolve("janji di tepati");
    }, 3000);
  } else {
    reject("janji gugur");
  }
});

async function getData() {
  try {
    const respone = await janji; //await adalah nunggu dulu
    console.log(respone);
  } catch (err) {
    console.log(err);
  }
}

getData();

// janji
//   .then((value) => {
//     console.log(value);
//   })
//   .catch((err) => {
//     console.log(err);
//   })
//   .finally(() => {
//     console.log("selesai");
//   });

// ---------------------------------------------------------------
// cara menggunakan sweet alert (untuk menampilkan error kepada user)
// cara menggunkana nya adalah input script cdn nya di html
//<script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
//lalu panggil saja code atau inputkan contoh2 nya di js ini

// let condition = false; // jadikan kondisi menjadi false biar masuk error

// let janji = new Promise((resolve, reject) => {
//   if (condition) {
//     setTimeout(() => {
//       resolve("janji di tepati");
//     }, 3000);
//   } else {
//     reject("janji gugur");
//   }
// });

// async function getData() {
//   try {
//     const respone = await janji; //await adalah nunggu dulu
//     console.log(respone);
//     Swal.fire("Good job!", respone, "success");
//   } catch (err) {
//     Swal.fire({
//       icon: "error",
//       title: "Oops...",
//       text: "regis gagal!",
//     });
//     console.log(err);
//   }
// }

// getData(); // jika menggukan sweet alert func ini harus di panggil menggunakan tombol button di halaman html nya

// promise adalah sebuah interface (seperti oop)
// tips (jika ada kata interface , class , new namaclass, itu adalah milik oop)
// status promise : pending , fullfilled (resolved), rejected
// promise digunakan ketika kita mau me request ke internet membutuhkan proses yg namanya asynchronous

let condition = false;

let janji = new Promise((resolve, reject) => {
  if (condition) {
    resolve("janji di tepati");
  } else {
    reject("janji gugur");
  }
});

//then untuk menangkap yang resolve
// catch untuk menangkap yang rejected
// finally mau hasil then catch tetap di jalankan
janji
  .then((value) => {
    console.log(value);
  })
  .catch((err) => {
    console.log(err);
  })
  .finally(() => {
    console.log("selesai");
  });

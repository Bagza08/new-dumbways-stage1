//asynchronous concept (menggambarkan asyncronus bekerja)
// dalam waktu bersama settimeout contoh async seakan2 mengambil data dari server dan menunggu selama 5 detik

console.log("hallo world");

setTimeout(() => {
  return console.log("1-javascript");
}, 5000); // 5000ms --> 5 Detik

setTimeout(() => {
  return console.log("2-javascript");
}, 7000); // 7000ms --> 7 Detik

setTimeout(() => {
  return console.log("3-javascript");
}, 5000); // 5000ms --> 5 Detik

console.log("Coder");

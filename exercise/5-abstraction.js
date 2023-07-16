// menampilkan data yang penting saja kepada client
// tujuan nya menyembunyikan detail yang gak penting dan user gak perlu tau

class Car {
  #name = "";
  #model = "";
  #duit = 0;

  constructor(name, model, duit) {
    this.#name = name;
    this.#model = model;
    this.#duit = duit;
  }

  //   getter;
  get name() {
    return this.#name;
  }

  get model() {
    return this.#model;
  }

  get duit() {
    return this.#duit;
  }
}

let myCar = new Car("toyota", "xy", 15000);

console.log(myCar.duit);
// console.log(myCar.model);

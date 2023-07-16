// contoh class concept / class adalah cetakan

class car {
  // properties = warna , jmlh roda, harga
  // method = autoDrive, gas , rem
  //   constructor adalah key untuk menampung nilai
  constructor(color, price) {
    this.color = color;
    this.price = price;
  }

  //   bikin method getinfo
  getInfo() {
    return `i have car with color ${this.color}, i buy it in ${this.price}`;
  }
}

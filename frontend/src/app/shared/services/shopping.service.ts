import { Injectable } from '@angular/core';

import { Product } from '../models/Product';

@Injectable()
export class ShoppingService {
  cartItems: Object[];

  constructor() {
    this.cartItems = [];
  }

  addItemToCart(p: Product) {
    this.cartItems.push(p);
  }
}
import { Injectable } from '@angular/core';

import { Product } from './Product';

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
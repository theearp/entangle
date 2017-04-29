import { Injectable } from '@angular/core';

import { Listing } from './Product';

@Injectable()
export class ShoppingService {
  cartItems: Object[];

  constructor() {
    this.cartItems = [];
  }

  addItemToCart(p: Listing) {
    this.cartItems.push(p);
  }
}
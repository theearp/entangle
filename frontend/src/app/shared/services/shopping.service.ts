import { Injectable } from '@angular/core';

import { Listing } from '../models/Listing';

@Injectable()
export class ShoppingService {
  cartItems: Object[];

  constructor() {
    this.cartItems = [];
  }

  addItemToCart(l: Listing) {
    this.cartItems.push(l);
  }
}
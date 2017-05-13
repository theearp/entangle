import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { BehaviorSubject } from 'rxjs';

import { Listing } from '../models/listing';
import { Cart } from '../models/cart';

import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/publishReplay';

@Injectable()
export class ShoppingService {
  private shoppingCart: BehaviorSubject<Listing[]> = new BehaviorSubject([]);
  cartItems: Listing[];

  constructor() {
    this.shoppingCart.subscribe(_ => this.cartItems = _);
  }

  add(l: Listing) {
    this.shoppingCart.next([...this.cartItems, l]);
  }

  get(): Observable<Listing[]> {
    return this.shoppingCart;
  }
}
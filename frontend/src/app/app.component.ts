import { Component, OnInit } from '@angular/core';

import { ShoppingService } from './shared/services/shopping.service';
import { CategoryService } from './shared/services/category.service';
import { Category } from './shared/models/category';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  categories: Category[];

  constructor(private ss: ShoppingService, private cs: CategoryService) {}

  showCart() {
    console.log('Number of items in cart: ' + this.ss.cartItems.length);
  }

  ngOnInit() {
    this.cs.getCategories()
     .subscribe(data => this.categories = data);
  }
}

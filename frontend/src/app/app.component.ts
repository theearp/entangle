import { Component, OnInit } from '@angular/core';

import { ShoppingService } from './shared/services/shopping.service';
import { CategoryService } from './shared/services/category.service';
import { Category, Section } from './shared/models/category';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  categories: Category[];
  sections: Section[];
  cartCount: Number;

  constructor(private ss: ShoppingService, private cs: CategoryService) {}

  showCart() {
    console.log('Number of items in cart: ' + this.cartCount);
  }

  ngOnInit() {
     this.cs.getSections()
      .subscribe(data => this.sections = data);

      this.ss.get()
      .subscribe(data => {
        this.cartCount = data.length;
      });
  }
}

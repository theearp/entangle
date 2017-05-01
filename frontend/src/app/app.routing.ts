import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ProductComponent } from './product/product.component';
import { ProductDetailComponent } from './product/productDetail.component';
import { PopularProductsComponent } from './product/popular.component';
import { ProductCategoryComponent } from './product/productCategory.component';

const routes: Routes = [
  { path: '', redirectTo: '/', pathMatch: 'full' },
  { path: 'product', component: ProductComponent},
  { path: 'product/details/:id', component: ProductDetailComponent},
  { path: 'popular', component: PopularProductsComponent },
  { path: 'product/category/:id', component: ProductCategoryComponent }
];
@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
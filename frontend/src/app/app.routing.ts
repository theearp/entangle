import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { ProductComponent } from './product.component';
import { ProductDetailComponent } from './productDetail.component';
import { ListingComponent } from './listing/listing.component';

const routes: Routes = [
  { path: '', redirectTo: '/', pathMatch: 'full' },
  { path: 'product', component: ProductComponent},
  { path: 'product/detail/:id', component: ProductDetailComponent},
  { path: 'listings', component: ListingComponent}
];
@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
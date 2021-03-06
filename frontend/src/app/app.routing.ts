import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';


import { AboutComponent } from './about/about.component';
import { ListingComponent} from './listing/listing.component';
import { ListingDetailComponent } from './listing/listing-detail.component';
import { CartComponent } from './cart/cart.component';

const routes: Routes = [
  { path: '', redirectTo: '/', pathMatch: 'full' },
  { path: 'about', component: AboutComponent},
  { path: 'listing', component: ListingComponent},
  { path: 'listing/:id', component: ListingDetailComponent},
  { path: 'shopping_cart', component: CartComponent}
];
@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
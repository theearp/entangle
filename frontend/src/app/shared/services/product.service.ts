import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/publishReplay';

import { Product, FakeProduct } from '../models/product';

export const PRODUCTS: FakeProduct[] = [
  {
    id: 1234, 
    name: 'Safe Snap Bullets', 
    img: '/assets/bullets.jpg', 
    category: 'Safe Snap', 
    description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
        Id mihi magnum videtur. A primo, ut opinor, animantium ortu petitur origo summi boni. 
        Duo Reges: constructio interrete. Stoicos roga. Quae cum magnifice primo dici viderentur, considerata minus probabantur. 
        Hoc unum Aristo tenuit: praeter vitia atque virtutes negavit rem esse ullam aut fugiendam aut expetendam. 
        Quid igitur dubitamus in tota eius natura quaerere quid sit effectum? Duo enim genera quae erant, fecit tria.`,
    price: '$10.99'
  },
  {
    id: 2345, 
    name: 'Very Cool Earrings', 
    img: '/assets/bullets.jpg', 
    category: 'Jewelry', 
    description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
        Id mihi magnum videtur. A primo, ut opinor, animantium ortu petitur origo summi boni. 
        Duo Reges: constructio interrete. Stoicos roga. Quae cum magnifice primo dici viderentur, considerata minus probabantur. 
        Hoc unum Aristo tenuit: praeter vitia atque virtutes negavit rem esse ullam aut fugiendam aut expetendam. 
        Quid igitur dubitamus in tota eius natura quaerere quid sit effectum? Duo enim genera quae erant, fecit tria.`,
    price: '$100.99'
  },
  {
    id: 3456, 
    name: 'Some Other Cool Product', 
    img: '/assets/bullets.jpg', 
    category: 'Unique', 
    description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
        Id mihi magnum videtur. A primo, ut opinor, animantium ortu petitur origo summi boni. 
        Duo Reges: constructio interrete. Stoicos roga. Quae cum magnifice primo dici viderentur, considerata minus probabantur. 
        Hoc unum Aristo tenuit: praeter vitia atque virtutes negavit rem esse ullam aut fugiendam aut expetendam. 
        Quid igitur dubitamus in tota eius natura quaerere quid sit effectum? Duo enim genera quae erant, fecit tria.`,
    price: '$34.99'
  },
  {
    id: 4567, 
    name: 'Wonder bread', 
    img: '/assets/bullets.jpg', 
    category: 'Bread', 
    description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
        Id mihi magnum videtur. A primo, ut opinor, animantium ortu petitur origo summi boni. 
        Duo Reges: constructio interrete. Stoicos roga. Quae cum magnifice primo dici viderentur, considerata minus probabantur. 
        Hoc unum Aristo tenuit: praeter vitia atque virtutes negavit rem esse ullam aut fugiendam aut expetendam. 
        Quid igitur dubitamus in tota eius natura quaerere quid sit effectum? Duo enim genera quae erant, fecit tria.`,
    price: '$346.99'
  }
];


@Injectable()
export class ProductService {
  baseUrl = 'http://localhost:8181/';
  allProducts = this.baseUrl + 'products';
  popularProducts = this.baseUrl + 'popular';
  productByID = this.baseUrl + 'product'
  productByCategory = this.baseUrl + 'product_category';

  constructor(private http: Http) {}

  // getProducts returns an observable of static products above after a 1
  // second delay.
  getFakeProducts(): Observable<Product[]> {
    return new Observable(observer => {
      setTimeout(() => {
        observer.next(PRODUCTS);
      }, 1000);
       setTimeout(() => {
        observer.complete();
       }, 2000);
    });
  };

  // getProduct fetchs a product by id from the static products above.
  getFakeProduct(id: number): Observable<Product> {
    return new Observable(observer => {
      PRODUCTS.forEach(function(p, i) {
        if (p.id == id) {
          observer.next(p);
        }
      });
      observer.complete();
    });
  };

  // getProducts fetches product from the GoApi -> CloudSQL data.
  getProducts(): Observable<Product[]> {
    return this.http.get(this.allProducts)
     .map(response => <Product[]> response.json())
     .catch(this.handleError);
  }

  getPopularProducts(): Observable<Product[]> {
    return this.http.get(this.popularProducts)
    .map(response => <Product[]> response.json())
    .catch(this.handleError);
  }

  getProduct(id): Observable<Product> {
    return this.http.get(this.productByID + '/' + id)
      .map(response => <Product> response.json())
      .catch(this.handleError)
      .publishReplay(1)
      .refCount()
  }

  getProductsByCategory(id): Observable<Product[]> {
    return this.http.get(this.productByCategory + '/' + id)
    .map(response => <Product> response.json())
      .catch(this.handleError)
      .publishReplay(1)
      .refCount()
  }

  private handleError (error: Response | any) {
    let errMsg: string;
    if (error instanceof Response) {
      const body = error.json() || '';
      const err = body.error || JSON.stringify(body);
      errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
    } else {
      errMsg = error.message ? error.message : error.toString();
    }
    console.error(errMsg);
    return Observable.throw(errMsg);
  }
};
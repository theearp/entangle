import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { Product, Listing } from './product';

export const PRODUCTS: Product[] = [
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

  constructor(private http: Http) {}

  getProducts(): Promise<Product[]> {
    return Promise.resolve(PRODUCTS);
  };

  getProduct(id: number) {
    let result: Product;
    PRODUCTS.forEach(function(p, i) {
      if (p.id == id) {
        result = p;
      }
    })
    if (result !== null){
      return Promise.resolve(result);
    } else {
      return Promise.reject(id + ' id not found in DB');
    }
  };

  getListings(): Observable<Listing[]> {
    return this.http.get('http://localhost:8181/products')
     .map(this.extractData)
     .catch(this.handleError);
  }

  private extractData(res: Response) {
    let body = res.json();
    return body.data || [];
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
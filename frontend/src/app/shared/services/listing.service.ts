import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';

import { Listing } from '../models/listing';
import { Image } from '../models/image';

import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/publishReplay';

@Injectable()
export class ListingService {
  baseUrl = 'http://localhost:8181/';
  listings = this.baseUrl + 'listings';
  listing = this.baseUrl + 'listing/';


  constructor(private http: Http) {}

  getListings(): Observable<Listing[]> {
    return this.http.get(this.listings)
     .map(response => <Listing[]> response.json() || [])
     .catch(this.handleError);
  }

  getListing(id: string): Observable<Listing> {
    return this.http.get(this.listing + id)
    .map(response => <Listing> response.json() || {})
    .catch(this.handleError);
  }

  getImages(id: string): Observable<Image[]> {
    return this.http.get(this.listing + id + '/images')
    .map(data => <Image[]> data.json() || [])
    .catch(this.handleError);
  }

  syncListing(id: string): Observable<String> {
    return this.http.get(this.baseUrl + '/listing/' + id + '/sync')
    .map(resp => resp.json() || '')
    .catch(this.handleError);
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
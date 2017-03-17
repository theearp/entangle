import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  recs = [
    {'name': 'Safe Snap Bullets', 'img': '/assets/bullets.jpg'},
    {'name': 'Very Cool Earrings', 'img': '/assets/bullets.jpg'},
    {'name': 'Some Other Cool Product', 'img': '/assets/bullets.jpg'},
    {'name': 'One more product', 'img': '/assets/bullets.jpg'},
  ]
  recCols: number;
  
  constructor() {
    this.recCols = this.recs.length;
  }
}

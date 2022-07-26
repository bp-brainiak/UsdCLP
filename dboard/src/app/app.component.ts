import { Component,OnInit, OnDestroy } from '@angular/core';
import { UsdSpotService } from './usdspot.service';
import { LabelRotation, LineStyle } from '@progress/kendo-angular-charts';
import { interval, Subscription  } from 'rxjs';
import { ThemeService } from '@progress/kendo-angular-charts/common/theme.service';
import { FormControl } from '@angular/forms';
import { MatCheckboxChange } from '@angular/material/checkbox';

import { ButtonFillMode, ButtonRounded, ButtonSize, ButtonThemeColor } from '@progress/kendo-angular-buttons';

export interface chartData  {
  timestamp: string;
  low:number;
  high:number;
  open:number;
  close: number;
  avg_01: number;
}


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [ UsdSpotService ],
})
export class AppComponent {


  // run with  docker run --rm -it -p 8999:80 angular-nginx
  constructor(private usdSpotService: UsdSpotService) {}

  subscription!: Subscription;
  chartsub!: Subscription
  source = interval(15000);
  source2 = interval(15000);
  public style: LineStyle = "smooth"
  public rotation: LabelRotation = { angle:270};
  cCount: number =1;
  bid:number=0;
  ask:number=0;
  spread:number=0;
  xrate:number=0;
  uf:string='';
  ipc:string='';
  utm:string=''
  status:string = '';
  chartStatus:string= '';
  range: string = '5d';
  region: string = 'US';
  interval: string = '1d';
  lang: string = 'en';
  recargaAuto: boolean = false;
  previousClose: number=0;
  marketState: string = '';
  cModel: chartData[] = [];
  cModelR: chartData[] = [];

  public size: ButtonSize = "large";
  public rounded: ButtonRounded = "large";
  public fillMode: ButtonFillMode = "solid";
  public themeColor: ButtonThemeColor = "success";

  ngOnInit(){

    if(localStorage.getItem("rango") != null) {
      this.range = localStorage.getItem("rango")!;
    }
    if(localStorage.getItem("rango") != null) {
      this.interval = localStorage.getItem("intervalo")!
    }

    this.getDollarSpot();
    this.getChartSpot();

    this.chartStatus='Disabled';
    this.status = 'Disabled';
    this.getUf();
    this.getIpc();
    this.getUtm();
  }

  OnDestroy() {
    this.stopSpotInterval();
    this.stopChartInterval();
  }

  onChangeDemo(ob: MatCheckboxChange) {
    console.log("checked: " + ob.checked);
    if(this.recargaAuto) {
      this.startSpotInterval();
      this.startChartInterval();
      this.chartStatus='Ok.';
      this.status = 'ok..';
    }
    else {
      this.chartStatus='Disabled';
      this.status = 'Disabled';
      this.stopSpotInterval();
      this.stopChartInterval();
    }
 }

  public onButtonClick(): void {
    this.getDollarSpot();
    this.getChartSpot();
  }
  padZero(num:number, size:number): string {
    let s = num+"";
    while (s.length < size) s = "0" + s;
    return s;
  }

  startSpotInterval(){
    this.subscription = this.source.subscribe(val=> this.getDollarSpot());
  }
  startChartInterval(){
    this.chartsub = this.source2.subscribe(val=> this.getChartSpot());
  }
  stopSpotInterval() {
    this.subscription.unsubscribe();
  }
  stopChartInterval() {
    this.chartsub.unsubscribe();
  }
  getDollarSpot()
  {
    localStorage.setItem("rango",this.range);
    localStorage.setItem("intervalo", this.interval);

    this.status = 'requesting..';
    this.usdSpotService.getDollar().subscribe((data: any) => {
      this.status = 'procesing...';
      this.processDataAfterRequest(data);
      if(!this.recargaAuto) {
        this.status = 'Disabled';
      }
    });

  }

  getUf() {
    this.uf= '';
    this.usdSpotService.getUfFromSbif().subscribe((data:any) =>{
      this.uf=data.UFs[0].Valor; // + ' al dia ' + data.UFs[0].Fecha
    });
  }

  getUtm(){
    this.utm= '';
    this.usdSpotService.getUtmFromSbif().subscribe((data:any) =>{
      this.utm=data.UTMs[0].Valor; // + ' al dia ' + data.UFs[0].Fecha
    });
  }

  getIpc(){
    this.ipc= '';
    this.usdSpotService.getIpcFromSbif().subscribe((data:any) =>{
      this.ipc=data.IPCs[0].Valor; // + ' al dia ' + data.UFs[0].Fecha
    });
  }
  
  processDataAfterRequest(data: any){
   if(data.quoteResponse.result.length == this.cCount)
   {
      this.status = 'ok...';
      this.bid = data.quoteResponse.result[0].bid;
      this.ask = data.quoteResponse.result[0].ask;
      this.marketState = data.quoteResponse.result[0].marketState
      this.xrate = (this.bid + this.ask) /2;
      this.spread = this.ask - this.bid;
   }
   else
   {
    this.status = 'different...';
   }

  }

  getChartSpot()
  {
    this.chartStatus='Requesting...';
    this.usdSpotService.getChartFromBackend(this.range, this.region, this.interval, this.lang)
      .subscribe((data: any)=> {
          this.processChartData(data);
       }
    );
  }
  processChartData(data: any){
    this.chartStatus='Procesando...'
console.log(data)
    try {
      var quote = data.chart.result[0].indicators.quote[0]
      this.previousClose =  data.chart.result[0].meta.chartPreviousClose;
      this.cModel = [];
      this.cModelR = [];
      console.log(data.chart.result[0].timestamp.length);
      for(let i=0; i <data.chart.result[0].timestamp.length; i++) {
        var date = new Date(data.chart.result[0].timestamp[i] * 1000);
  
          this.cModel.push({
            "timestamp": this.padZero(date.getDate(),2) + "-" + this.padZero(date.getMonth() + 1, 2)+ "-" + date.getFullYear() + " " +  this.padZero(date.getHours(),2) + ":" + this.padZero(date.getMinutes(),2),
            "low": quote.low[i],
            "high": quote.high[i],
            "open": quote.open[i],
            "close": quote.close[i],
            "avg_01": (quote.open[i]+quote.close[i] + quote.low[i] +quote.high[i])  /4
          }); 
 
      }
      this.clearNullData();
      this.chartStatus='Ok..'
      for(let j = 0; j < this.cModel.length; j++){
        this.cModelR.push(this.cModel[j]);
      }
      this.cModelR = this.cModelR.reverse();
    }
    catch(error){
      this.chartStatus='error procesing request..'
    }


    if(!this.recargaAuto) {
      this.chartStatus = 'Disabled';
    }
  }

  private clearNullData() {
    var cModel2: chartData[] = [];
    for(var i = 0; i <this.cModel.length; i++) {
      if (this.cModel[i].low == null || this.cModel[i].high == null ||this.cModel[i].open == null || this.cModel[i].close == null) {
        continue;
      }
      cModel2.push(this.cModel[i])
    }
    this.cModel = cModel2;
  }
}

import { Injectable, } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { environment } from "src/environments/environment";
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators'; 


@Injectable() 
export class ConfigService {

    constructor(private http: HttpClient) {}
    private httpOptions = {
         headers: new HttpHeaders()
         .set('Content-Type', 'application/json')
         .set('X-API-KEY','NjU4UkMogG8ad1tthie626TYkVcIaSc11YuEwtk6')
    };

    getDollar(): Observable<any> {
        console.log(environment.bakendHost+"/usdSpot");
        return this.http.get<any>(environment.bakendHost+"/usdSpot");
    }

    getChartData(range: string, region: string, interval: string, lang:string ): Observable<any> {
        //range=1mo&region=US&interval=1d&lang=en&events=div%2Csplit'
        var _range =`range=${range}`;
        var _region =`region=${region}`;
        var _interval =`interval=${interval}`;
        var _lang =`lang=${lang}`;
        var endpoint = `${environment.chartEndpoint}${_range}&${_region}&${_interval}&${_lang}`
        return this.http.get<any>(endpoint, this.httpOptions);
    }
}
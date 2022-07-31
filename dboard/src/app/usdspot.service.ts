import { Injectable, } from "@angular/core";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { environment } from "src/environments/environment";
import { Observable, throwError } from 'rxjs';
import { catchError, retry } from 'rxjs/operators'; 
import { ThemeService } from "@progress/kendo-angular-charts/common/theme.service";


@Injectable() 
export class UsdSpotService {

    constructor(private http: HttpClient) {}
    private httpOptions = {
         headers: new HttpHeaders()
         .set('Content-Type', 'application/json')
         .set('X-API-KEY','')
    };

    getDollar(): Observable<any> {
        return this.http.get<any>(environment.bakendHost+"/usdSpot");
    }
    getChartFromBackend(range: string, region: string, interval: string, lang:string): Observable<any> {
        var _range =`range=${range}`;
        var _region =`region=${region}`;
        var _interval =`interval=${interval}`;
        var _lang =`lang=${lang}`;
        var endpoint = `${environment.bakendHost}/chart?${_range}&${_region}&${_interval}&${_lang}`;
        console.log(endpoint);
        return this.http.get<any>(endpoint);
    }
    getUfFromSbif():Observable<any> {
        return this.http.get<any>(environment.endPointUF);
    }
}    

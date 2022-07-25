import { NgModule } from "@angular/core"; 
import { MatToolbarModule } from "@angular/material/toolbar"; 
import { MatTabsModule } from '@angular/material/tabs'; 
import { MatCardModule } from '@angular/material/card'; 
import { MatButtonToggleModule } from '@angular/material/button-toggle'; 
import { MatGridListModule } from '@angular/material/grid-list'; 
import { MatRadioModule } from '@angular/material/radio'; 
import { MatCheckboxModule } from '@angular/material/checkbox'; 
@NgModule({ 
    exports: [ 
        MatToolbarModule, 
        MatTabsModule,
        MatCardModule,
        MatButtonToggleModule,
        MatGridListModule,
        MatRadioModule,
        MatCheckboxModule,
        MatCheckboxModule
    ]
})
export class MaterialModule {}
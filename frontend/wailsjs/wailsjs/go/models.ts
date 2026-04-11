export namespace main {
	
	export class Transaction {
	    id: number;
	    product: string;
	    applicant_name: string;
	    date: string;
	    purpose: string;
	    explanation: string;
	    price: number;
	    expenditure: string;
	    category: string;
	    status: string;
	    rejected_reason: string;
	    receipt: string;
	    is_advance_application: string;
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.product = source["product"];
	        this.applicant_name = source["applicant_name"];
	        this.date = source["date"];
	        this.purpose = source["purpose"];
	        this.explanation = source["explanation"];
	        this.price = source["price"];
	        this.expenditure = source["expenditure"];
	        this.category = source["category"];
	        this.status = source["status"];
	        this.rejected_reason = source["rejected_reason"];
	        this.receipt = source["receipt"];
	        this.is_advance_application = source["is_advance_application"];
	    }
	}

}


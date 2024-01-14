export interface Employment {
    role: string;
    user: User;
    business: Business;
}

export interface User {
    uid: string;
    phone: string | null;
}

export interface Business {
    uid: string;
    name: string;
}
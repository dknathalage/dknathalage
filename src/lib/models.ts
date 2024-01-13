export interface BusinessMember {
    role: string;
    user: User;
    business: Business;
}

export interface User {
    phone: string | null;
}

export interface Business {
    name: string;
}
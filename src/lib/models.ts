interface BusinessMember {
    uid: string;
    role: string;
    user: User;
    business: Business;
}

export interface User{
    uid: string;
    phone: string | null;
    employments: BusinessMember[];
}

export interface Business {
    uid: string;
    name: string;
    employments: BusinessMember[];
}
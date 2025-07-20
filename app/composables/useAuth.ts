import type { User } from "~/types/types"

export const useAuth = () => {
    const user = useState<User | null>("user", () => null)
    const token = useState<string | null>("token", () => null)
    const loading = useState<boolean>("loading", () => true)
    
    const isAuthenticated = computed(() => !!user.value && !!token.value)
    return {
        user,
        token,
        loading,
        isAuthenticated
    } 
}

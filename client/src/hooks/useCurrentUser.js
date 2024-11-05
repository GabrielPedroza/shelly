import { useQuery } from '@tanstack/react-query';

async function fetchUser() {
	const response = await fetch('http://localhost:8000/api/user', {
		credentials: 'include',
	});

	if (!response.ok) {
		throw new Error('Failed to fetch user session');
	}

	return response.json();
}

export default function useCurrentUser() {
	return useQuery({
		queryKey: ['currentUser'],
		queryFn: fetchUser,
		retry: false,
		staleTime: 1000 * 60 * 5,
		onError: (error) => console.error('Error fetching user:', error),
	});
}

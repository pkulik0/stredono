
export interface EmailMessage {
	id: number
	from: string
	subject: string
	date: string
}

export const getMessages = async (email: string): Promise<EmailMessage[]> => {
	const login = email.split('@')[0]
	const domain = email.split('@')[1]

	const res = await fetch(`https://www.1secmail.com/api/v1/?action=getMessages&login=${login}&domain=${domain}`)
	if(!res.ok) {
		throw new Error('Failed to get messages')
	}

	return await res.json() as EmailMessage[]
}

export const getMessage = async (email: string, id: number): Promise<string> => {
	const login = email.split('@')[0]
	const domain = email.split('@')[1]

	const res = await fetch(`https://www.1secmail.com/api/v1/?action=readMessage&login=${login}&domain=${domain}&id=${id}`)
	if(!res.ok) {
		throw new Error('Failed to get message')
	}

	const data = await res.json()
	return data.body
}
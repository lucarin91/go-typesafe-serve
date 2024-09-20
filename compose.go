package tss

func Compose2[T1 any, T2 any, T3 any, F1 ~func(T1) T2, F2 ~func(T2) T3](f2 F2, f1 F1) func(T1) T3 {
	return func(t T1) T3 {
		return f2(f1(t))
	}
}

func Compose3[T1 any, T2 any, T3 any, T4 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4](f3 F3, f2 F2, f1 F1) func(T1) T4 {
	return func(t T1) T4 {
		return f3(f2(f1(t)))
	}
}

func Compose4[T1 any, T2 any, T3 any, T4 any, T5 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5](f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T5 {
	return func(t T1) T5 {
		return f4(f3(f2(f1(t))))
	}
}

func Compose5[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5, F5 ~func(T5) T6](f5 F5, f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T6 {
	return func(t T1) T6 {
		return f5(f4(f3(f2(f1(t)))))
	}
}

func Compose6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5, F5 ~func(T5) T6, F6 ~func(T6) T7](f6 F6, f5 F5, f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T7 {
	return func(t T1) T7 {
		return f6(f5(f4(f3(f2(f1(t))))))
	}
}

func Compose7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5, F5 ~func(T5) T6, F6 ~func(T6) T7, F7 ~func(T7) T8](f7 F7, f6 F6, f5 F5, f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T8 {
	return func(t T1) T8 {
		return f7(f6(f5(f4(f3(f2(f1(t)))))))
	}
}

func Compose8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5, F5 ~func(T5) T6, F6 ~func(T6) T7, F7 ~func(T7) T8, F8 ~func(T8) T9](f8 F8, f7 F7, f6 F6, f5 F5, f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T9 {
	return func(t T1) T9 {
		return f8(f7(f6(f5(f4(f3(f2(f1(t))))))))
	}
}

func Compose9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5, F5 ~func(T5) T6, F6 ~func(T6) T7, F7 ~func(T7) T8, F8 ~func(T8) T9, F9 ~func(T9) T10](f9 F9, f8 F8, f7 F7, f6 F6, f5 F5, f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T10 {
	return func(t T1) T10 {
		return f9(f8(f7(f6(f5(f4(f3(f2(f1(t)))))))))
	}
}

func Compose10[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any, T10 any, T11 any, F1 ~func(T1) T2, F2 ~func(T2) T3, F3 ~func(T3) T4, F4 ~func(T4) T5, F5 ~func(T5) T6, F6 ~func(T6) T7, F7 ~func(T7) T8, F8 ~func(T8) T9, F9 ~func(T9) T10, F10 ~func(T10) T11](f10 F10, f9 F9, f8 F8, f7 F7, f6 F6, f5 F5, f4 F4, f3 F3, f2 F2, f1 F1) func(T1) T11 {
	return func(t T1) T11 {
		return f10(f9(f8(f7(f6(f5(f4(f3(f2(f1(t))))))))))
	}
}

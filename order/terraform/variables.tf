variable "aws_region" {
  description = "Region"
  type        = string
  default     = "ap-southeast-1"
}

variable "tfuser_arn" {
  description = "arn"
  type        = string
  default     = "arn:aws:iam::841162681430:group/tf-1"
}
